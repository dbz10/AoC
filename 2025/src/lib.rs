use std::{env, fs};

pub fn read_content() -> String {
    let target_file = &env::args().collect::<Vec<String>>()[1];
    fs::read_to_string(target_file).expect("Unable to read file.")
}
