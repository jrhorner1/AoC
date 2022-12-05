/* 
Advent of Code 2022
Day 1 - Rust
*/
use std::fs;

fn main() {
    let filepath = "../../input/1";
    let raw_input: String = fs::read_to_string(filepath)
        .expect("Should have been able to read the file");
    let input = raw_input.split("\n\n");
    for inp in input {
        println!("{}", inp                                          );
    }

    let mut calories = vec![];
    for elf in input {
        let snacks = elf.split("\n");
        let mut cals = 0;
        for snack in snacks {
            let snack: u32 = snack.trim().parse()
                .expect("Something went wrong with the conversion.");
            cals = cals + snack;
        };
        calories.push(cals)
    }
    println!("{}", calories[1]);
}
