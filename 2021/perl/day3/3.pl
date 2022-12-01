#!/usr/bin/perl
# Pass 1 or 2 as a cli argument to solve for part 1 or part 2

use v5.30;
use strict;
use warnings;
no warnings 'experimental';

my $inputfile = '../../input/3';
open(FH, '<', $inputfile) or die $!;

my @input = ();
while(<FH>){
    $_ =~ s/\R//g; # Remove newline chars
    push(@input, $_);
}

sub part1() {
    my @bit_counts = ();
    # populate array with zeros
    for(my $i = 0; $i < length($input[0]); $i++){
        push(@bit_counts, 0);
    }
    for(@input){
        my @byte = split('', $_);
        for(my $i = 0; $i <= $#byte; $i++){
            if($byte[$i] == "1") {
                $bit_counts[$i]++;
            }
        }
    }
    my $gamma_string = "";
    my $epsilon_string = "";
    for(@bit_counts){
        if($_ >= $#input/2){
            $gamma_string = $gamma_string . "1";
            $epsilon_string = $epsilon_string . "0";
        } else {
            $gamma_string = $gamma_string . "0";
            $epsilon_string = $epsilon_string . "1";
        }
    }
    my $gamma_rate = oct("0b$gamma_string");
    # unary NOT doesn't work because of system int size. 
    # my $epsilon_rate = ~$gamma_rate; 
    my $epsilon_rate = oct("0b$epsilon_string");
    return $gamma_rate * $epsilon_rate;
}

sub part2() {
    return 2
}

if($ARGV[0] == 1){
    print part1(), "\n";
} else {
    print part2(), "\n";
}
close(FH);