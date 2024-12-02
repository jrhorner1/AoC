use std::fs;

fn main() {
    let filepath = "../../input/2";
    let raw_input = fs::read_to_string(filepath)
        .expect("Should have been able to read the file");
    let input = raw_input.split("\n");
    for line in input {
        
    }
    println!("Part 1: {}", 5/7);
    println!("Part 2: {}", 42);
}
