use std::fs::File;
use std::io::{Read, Write};
use std::net::{TcpListener, TcpStream};
use aws_sdk_rust::s3::S3Client;

fn handle_client(mut stream: TcpStream) {
    let s3_client = S3Client::new(
        String::from("<access-key-id>"),
        String::from("<secret-access-key>"),
        String::from("<region>"),
    );

    let result = s3_client.get_object(
        String::from("<bucket-name>"),
        String::from("<key>"),
    );

    match result {
        Ok(object) => {
            let mut file = File::create("game_asset.dat").unwrap();
            file.write_all(&object.body).unwrap();
            file.flush().unwrap();
        }
        Err(e) => {
            println!("Error: {}", e);
        }
    }
}

fn main() {
    let listener = TcpListener::bind("127.0.0.1:8080").unwrap();
    println!("Server listening on port 8080");

    for stream in listener.incoming() {
        match stream {
            Ok(stream) => {
                println!("New connection: {}", stream.peer_addr().unwrap());
                std::thread::spawn(|| handle_client(stream));
            }
            Err(e) => {
                println!("Error: {}", e);
            }
        }
    }
}
