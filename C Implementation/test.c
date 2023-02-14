/*
Project:        Data Streaming Using Raspberry Pi Pico W
Team:           Reconfigurable Space Computing Lab
Advisor:        Dr. Aly
Members:        Alexander Ea
                Shahzman Saqib
                Lino Mercado-Esquivias
Description:    This code turns a Pico W into a TCP server. The server
                waits for a client to connect and then sends 10 data
                packets. After each packet, the server waits for the client
                to send the packet back before sending the next. The server
                compares if the data received matches the data sent and 
                prints the result.
Version:        v2
*/

#include <stdio.h>
#include <stdlib.h>
#include "pico/stdlib.h"
#include "crypto_aead.h"

int main(void){
    stdio_init_all();

    // variables
    // const unsigned char encryption_m[] = {0x00, 0x01, 0x02, 0x03};                
    // const unsigned char encryption_m[] = {0x61, 0x62, 0x63, 0x64};
    const unsigned char encryption_m[] = {0x61, 0x62, 0x63, 0x64, 0x65, 0x66, 0x67, 0x68};
    // const unsigned char encryption_m[] = {0x61, 0x62, 0x63, 0x64, 0x65, 0x66, 0x67, 0x68, 0x69, 0x6A, 0x6B, 0x6C, 0x6D, 0x6E, 0x6F, 0x70};

    // unsigned char decryption_c[] = {0xE0, 0xE1, 0x5D, 0x6E, 0xD5, 0x3B, 0xA0, 0xFD, 0x92, 0x36, 0x02, 0x9D}; 
    // unsigned char decryption_c[] = {0x26, 0x6C, 0xE7, 0x3F, 0xB0, 0x11, 0x82, 0x01, 0x3E, 0x8C, 0x59, 0x3A};  
    unsigned char decryption_c[] = {0x26,  0x6C,  0xE7, 0x3F, 0x1F, 0x07, 0x39, 0x59, 0x97,  0xEE,  0xB7, 0x42,  0x81, 0x12,  0xB4, 0x1F};
    // unsigned char decryption_c[] = {0x26,  0x6C,  0xE7, 0x3F, 0x1F, 0x07, 0x39, 0x59, 0x8A,  0xE6, 0x08, 0x30,  0xDF, 0x1A,  0xF2, 0x4E,  0xEF,  0xD1,  0xC3,  0x86, 0x38,  0x82,  0xA6,  0xB4}; 

    unsigned long long mlen = 16;                                // decryption plaintext length
    unsigned long long clen = mlen + 8;                         // encryption ciphertext length
    unsigned char encryption_c[clen];                           // encryption ciphertext
    unsigned char decryption_m[mlen];                           // decryption plaintext
    const unsigned char ad[] = {0x00};                       // associated data
    // const unsigned char ad[] = {0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0A, 0x0B, 0x0C, 0x0D, 0x0E, 0x0F, 0x10};                          // associated data
    unsigned long long adlen = 0;                      // associated data length
    const unsigned char *encryption_nsec;                       // encryption secret message number
    unsigned char *decryption_nsec;                             // decryption secret message number
    const unsigned char npub[] = {0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0A, 0x0B};  // public message number
    const unsigned char k[] = {0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0A, 0x0B, 0x0C, 0x0D, 0x0E, 0x0F}; // 128-bit key

    // countdown
    for(int i = 5; i > 0; i--){
        sleep_ms(1000);
        printf("%d\n",i);
    }

    // encrypt
    crypto_aead_encrypt(encryption_c, &clen, encryption_m, mlen, ad, adlen, encryption_nsec, npub, k);
    printf("Ciphertext = ");
    for (int i = 0; i < clen; i++){
        printf("%02X", encryption_c[i]);
    }
    printf("\n");

    // decrypt
    crypto_aead_decrypt(decryption_m, &mlen, decryption_nsec, decryption_c, clen, ad, adlen, npub, k);

    printf("Plaintext = ");
    for (int i = 0; i < mlen; i++){
        printf("%02X", decryption_m[i]);
    }
    printf("\n");

    printf("Start!");
    for(int i = 0; i < 500000; i++){
        crypto_aead_encrypt(encryption_c, &clen, encryption_m, mlen, ad, adlen, encryption_nsec, npub, k);
        crypto_aead_decrypt(decryption_m, &mlen, decryption_nsec, decryption_c, clen, ad, adlen, npub, k);
    }
    printf("Finish!");
    return 0;
}