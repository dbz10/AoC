use std::{env, fs};

pub fn read_content() -> String {
    let args: Vec<String> = env::args().collect();
    let target_file = &args[2];
    fs::read_to_string(target_file).expect("Unable to read file.")
}
