/* 
Advent of Code 2022
Day 1 - Rust
*/
use std::fs;

fn main() {
    let filepath = "../../input/1";
    let raw_input = fs::read_to_string(filepath)
        .expect("Should have been able to read the file");
    let input = raw_input.split("\n\n");
    let mut elves: Vec<u32> = Vec::new();
    for bag in input {
        let snacks = bag.split("\n");
        let mut total_calories = 0;
        for snack in snacks {
            let calories: u32 = match snack.parse() {
                Ok(calories) => calories,
                Err(_) => continue,
            };
            total_calories += calories
        }
        elves.push(total_calories)
    }
    elves.sort();
    let most_calories = elves.last().unwrap();
    println!("Part 1: {}", most_calories);
    let mut top_three_total = 0;
    let mut count = 0;
    for elf in elves.iter().rev() {
        count += 1;
        top_three_total += elf;
        if count == 3 {
            break
        }
    }
    println!("Part 2: {}", top_three_total)
}
