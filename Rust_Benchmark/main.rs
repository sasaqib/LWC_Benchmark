// Version4: Goal -> Ciphertext now working, goal is to start timing 4/8/16-bytes of message
// Current problem: none?
// February 7th, 2023
#![no_std]
#![no_main]

use cortex_m::delay;
use cortex_m::prelude::_embedded_hal_blocking_delay_DelayMs;
use rust_tinyJAMBU::tinyJAMBU::*;


use bsp::entry;
use defmt::*;
use defmt_rtt as _;
// use embedded_hal::digital::v2::OutputPin;
use panic_probe as _;

// Provide an alias for our BSP so we can switch targets quickly.
// Uncomment the BSP you included in Cargo.toml, the rest of the code does not need to change.
use rp_pico as bsp;
// use sparkfun_pro_micro_rp2040 as bsp;

use bsp::hal;

use bsp::hal::{
    clocks::{init_clocks_and_plls, Clock},
    pac,
    // sio::Sio,
    watchdog::Watchdog,
};

use fugit::RateExtU32;

use usb_device::{class_prelude::*, prelude::*};
use usbd_serial::SerialPort;

use bsp::hal::uart::{DataBits, StopBits, UartConfig};


fn test_encrypt_pico(clen: &mut u64, c: &mut [u32]) {
    let mut state: [u32; 4] = [0; 4];
    //let input = "abcd";
    let mut c: [u32; 80] = [0; 80];
    // let mut clen: u64 = 80;
    let mut mlen: usize = 4; //Change this
     let mut m1: [u8; 4] = [97, 98, 99, 100];
    //let mut m1: [u8; 8] = [97, 98, 99, 100, 101, 102, 103, 104];
     //let mut m1: [u8; 16] = [97, 98, 99, 100, 101, 102, 103, 104, 105, 106, 107, 108, 109, 110, 111, 112];
    
   
    // let mut mlen: usize = 8;
    // let mut mlen: usize = 16;

    let ad: [char; 80] = ['\x00'; 80];
    let adlen: usize = 0;
    let npub: [u8; 12] = [0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11];
    let nsec: [u8; 80] = [0; 80];
    let k: [u8; 16] = [0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15];
    //string2u8(input, &mut m);
    // println!("Running state_update with 1024 steps\n");
    state_update(&mut state, &k, 1024);
    for state_elem in state {
        // println!("{}", state_elem);
    }
    // println!("Running initialization\n");
    initialization(&k, &npub, &mut state);
    for state_elem in state {
        // println!("{}", state_elem);
    }
    // println!("Running encryption\n");
    crypto_aead_encrypt(&mut c, clen, &m1, mlen, &ad, adlen, &nsec, &npub, &k);
    let mut i: usize = 0;
    while i < (*clen * 4) as usize {
        // println!("{:#04X}", u32_to_u8(&c, i));
        i = i + 1;
    }
}
fn test_decrypt_pico(m: &mut [u32], mlen: &mut u64, status: &mut bool) {
    let mut state: [u32; 4] = [0; 4];
    let input = "\x00";
    //let mut c: [u8; 80] = [0; 80];
    let mut c_temp: [u8; 80] = [0; 80];
    let clen: u64 = 12; //Change this
    let mut c1: [u8; 12] = [38, 108, 231, 63, 176, 17, 130, 1, 62, 140, 89, 58]; //abcd
    //let mut c1: [u8; 16] = [38, 108, 231, 63, 31, 7, 57, 89, 151, 238, 183, 66, 129, 18, 180, 31]; //abcdefgh
    // let mut c1 :[u8; 24] = [38, 108, 231, 63, 31, 7, 57, 89, 138, 230, 8, 48, 223, 26, 242, 78, 239, 209, 195, 134, 56, 130, 166, 180]; //abcdefghijklmnop

    //string2u8("266CE73FB01182013E8C593A", &mut c_temp);
    // println!("Printing string as u8: ");
    let mut hex_char: [u8; 80] = [0; 80];
    for (ch, hex) in c_temp.iter().zip(hex_char.iter_mut()) {
        *hex = hex_char_to_decimal(*ch);
    }
    // println!("Printing converted to decimal: ");
    for ch in hex_char {
        // println!("{}", ch);
    }
    let mut combined_char: [u8; 80] = [0; 80];
    u4tou8(&hex_char, &mut combined_char);
    // println!("Printing combined char: ");
    for ch in combined_char {
        // println!("{}", ch);
    }
    for (combined, ch) in combined_char.iter_mut().zip(c1.iter_mut()) {
        *ch = *combined;
    }
    // println!("printing ciphertext");
    for ch in c1 {
        // println!("{}", ch);
    }
   
    // let mut m: [u32; 80] = [0; 80];
    // let mut mlen: u64 = 80;
    let ad: [char; 80] = ['\x00'; 80];
    let adlen: usize = 0;
    let npub: [u8; 12] = [0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11];
    let mut nsec: [u8; 80] = [0; 80];
    let k: [u8; 16] = [0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15];
    // println!("testing decryption");

    if (crypto_aead_decrypt(m, mlen, &mut nsec, &c1, clen, &ad, adlen, &npub, &k) == 0) {
        *status = true;
    } else {
        *status = false;
        // println!("decryption tag is not verified");
    }
}

pub fn convert_u32_arr(data: &[u32; 4]) -> [u8; 16] {
    let mut res = [0; 16];
    for i in 0..4 {
        res[4 * i..][..4].copy_from_slice(&data[i].to_be_bytes());
    }
    res
}

fn val_to_u8_arr(arr_out: &mut [u8], input: &mut u32) {
    let mut i = 0;
    while *input > 0 {
        arr_out[i] = char::from_digit(*input % 10, 10).unwrap() as u8;
        *input /= 10;
        i = i + 1;
    }
    arr_out.reverse();
}

// pub fn convert(source: &[u32; 4]) -> [u8; 16] {
//     let mut dest = [0; 16];
//     for (dest_c, source_e) in dest.chunks_exact_mut(4).zip(source.iter()) {
//         dest_c.copy_from_slice(&source_e.to_le_bytes())
//     }
//     dest
// }
pub fn convert(data: &[u32; 80]) -> [u8; 12] {
    let mut res = [0; 12];
    for i in 0..80 {
        res[80 * i..][..80].copy_from_slice(&data[i].to_ne_bytes());
    }
    res
}
fn val_to_u8_arr_cpy(arr_out: &mut [u8], input: &u32) {
    let mut i = 0;
    let mut input_temp = *input;
    while input_temp > 0 {
        arr_out[i] = char::from_digit(input_temp % 10, 10).unwrap() as u8;
        input_temp /= 10;
        i = i + 1;
    }
    arr_out.reverse();
}

fn test_write_buffer(UART_buffer: &mut [u8]) {
    let mut c: [u32; 80] = [0; 80];
    //let mut c: [u32; 12] = [0; 12];
    let mut clen: u64 = 80;
    test_encrypt_pico(&mut clen, &mut c);

    let mut i = 0;
    while i < (clen * 4) as usize {
        UART_buffer[i] = u32_to_u8(&c, i);
        i = i + 1;
    }
}

#[entry]
fn main() -> ! {
    info!("Program start");
    let mut pac = pac::Peripherals::take().unwrap();
    let core = pac::CorePeripherals::take().unwrap();
    let mut watchdog = Watchdog::new(pac.WATCHDOG);
    //let sio = Sio::new(pac.SIO);

    // External high-speed crystal on the pico board is 12Mhz
    let external_xtal_freq_hz = 12_000_000u32;
    let clocks = init_clocks_and_plls(
        external_xtal_freq_hz,
        pac.XOSC,
        pac.CLOCKS,
        pac.PLL_SYS,
        pac.PLL_USB,
        &mut pac.RESETS,
        &mut watchdog,
    )
    .ok()
    .unwrap();

    let mut delay = cortex_m::delay::Delay::new(core.SYST, clocks.system_clock.freq().to_Hz());

    // let pins = bsp::Pins::new(
    //     pac.IO_BANK0,
    //     pac.PADS_BANK0,
    //     sio.gpio_bank0,
    //     &mut pac.RESETS,
    // );

    let usb_bus = UsbBusAllocator::new(hal::usb::UsbBus::new(
        pac.USBCTRL_REGS,
        pac.USBCTRL_DPRAM,
        clocks.usb_clock,
        true,
        &mut pac.RESETS,
    ));

    let mut serial = SerialPort::new(&usb_bus);
    let mut usb_dev = UsbDeviceBuilder::new(&usb_bus, UsbVidPid(0x16c0, 0x27dd))
        .manufacturer("Fake company")
        .product("Serial port")
        .serial_number("TEST")
        .device_class(2) // from: https://www.usb.org/defined-class-codes
        .build();

    let sio = hal::Sio::new(pac.SIO);

    // let mut led_pin = pins.led.into_push_pull_output();
    let pins = hal::gpio::Pins::new(
        pac.IO_BANK0,
        pac.PADS_BANK0,
        sio.gpio_bank0,
        &mut pac.RESETS,
    );

    let uart_pins = (
        // UART TX (characters sent from RP2040) on pin 1 (GPIO0)
        pins.gpio8.into_mode::<hal::gpio::FunctionUart>(),
        // UART RX (characters received by RP2040) on pin 2 (GPIO1)
        pins.gpio9.into_mode::<hal::gpio::FunctionUart>(),
    );

    let mut uart = hal::uart::UartPeripheral::new(pac.UART1, uart_pins, &mut pac.RESETS)
        .enable(UartConfig::default(), clocks.peripheral_clock.freq())
        .unwrap();

    let timer = hal::Timer::new(pac.TIMER, &mut pac.RESETS);
    let mut current_time = timer.get_counter();
    loop {
        if usb_dev.poll(&mut [&mut serial]) {}
        if timer.get_counter() - current_time >= 2_000_000 {
            // let _ = serial.write(b"UART_Write_start \r\n");
            // let mut temp_buf: [u8; 80] = [0; 80];
            // test_write_buffer(&mut temp_buf);
            // uart.write_full_blocking(&temp_buf);
                delay.delay_ms(500);
                let _ = serial.write(b"1\r\n");
                delay.delay_ms(500);
                let _ = serial.write(b"2\r\n");
                delay.delay_ms(500);
                let _ = serial.write(b"3\r\n");
                delay.delay_ms(500);
                let _ = serial.write(b"4\r\n");
                delay.delay_ms(500);
                let _ = serial.write(b"5\r\n");
                delay.delay_ms(500);

            
                let mut clen: u64 = 12; // change this
                let mut c: [u32;80] = [0;80];
                let mut mlen: u64 = 4; // change this
                let mut m: [u32;80] = [0;80];
                let mut status: bool = false;
                let _ = serial.write(b"Start Timing for 4-bytes\r\n"); // change this

            for i in 0..500000{
                test_encrypt_pico(&mut clen, &mut c);
                test_decrypt_pico(&mut m, &mut mlen, &mut status);
            }
            let _ = serial.write(b"Program Finished\r\n");
            
        }
    }
    
}
// let mut read_buf: [u8; 80] = [0; 80];

            // let _ = serial.write(b"UART_Read_start \r\n");
            // let _ = uart.read_full_blocking(&mut read_buf);
            // let _ = serial.write(&mut read_buf);
// End of file

// CT: 38 108 231 63 176 17 130 1 62 140 89 58
// 3810823163176171301621408958 0536927692

// 1. Server
// 2. UART_Write
// 3. Client
// 4. UART_Read


// Timings:

// Round1:
// 04-byte: 2:00.62
// 08-byte: 2:13.36
// 16-byte: 2:38.86

// Round2:
// 04-byte: 2:00.61
// 08-byte: 2:13.28
// 16-byte: 2:38.71

// Round3:
// 04-byte: 2:00.67
// 08-byte: 2:13.51
// 16-byte: 2:38.82