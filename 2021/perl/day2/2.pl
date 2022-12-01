#!/usr/bin/perl
# Pass 1 or 2 as a cli argument to solve for part 1 or part 2

use v5.10;
use strict;
use warnings;
no warnings 'experimental';

my $inputfile = '../../input/2';
open(FH, '<', $inputfile) or die $!;

my @input = ();
while(<FH>){
    push(@input, $_)
}

sub part1() {
    my $horizontal = 0;
    my $depth = 0;
    for(@input) {
        my @line = split(' ', $_);
        my $dir = $line[0];
        my $dist = $line[1];
        given($dir){
            when(/forward/) { $horizontal += $dist }
            when(/down/) { $depth += $dist }
            when(/up/) { $depth -= $dist }
        }
    }
    return $horizontal * $depth;
}

sub part2() {
    my $horizontal = 0;
    my $depth = 0;
    my $aim = 0;
    for(@input) {
        my @line = split(' ', $_);
        my $dir = $line[0];
        my $dist = $line[1];
        given($dir){
            when(/forward/) { 
                $horizontal += $dist;
                $depth += $aim * $dist;
            }
            when(/down/) { $aim += $dist }
            when(/up/) { $aim -= $dist }
        }
    }
    return $horizontal * $depth;
}

if($ARGV[0] == 1){
    print part1(), "\n";
} else {
    print part2(), "\n";
}
close(FH);