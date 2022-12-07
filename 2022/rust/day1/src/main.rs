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
    let top_three_total = most_calories + elves[elves.len() - 2] + elves[elves.len() - 3];
    println!("Part 2: {}", top_three_total)
}
