/* 
Advent of Code 2022
Day 1 - Rust
*/
use std::fs;

fn main() {
    let filepath = "../../input/1";
    let raw_input = fs::read_to_string(filepath)
        .expect("should have been able to read the file");
    let input = raw_input.split("\n\n");
    let mut calories = vec![];
    for elf in input {
        let snacks = elf.split("\n");
        let mut cals = 0;
        for snack in snacks {
            cals += snack.parse::<u32>("\n").ok();
        }
        calories.push(cals)
    }
    println!("{}", calories[1]);
}