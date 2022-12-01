#!/usr/bin/perl
# Pass 1 or 2 as a cli argument to solve for part 1 or part 2

use v5.30;
use strict;
use warnings;
no warnings 'experimental';

my $inputfile = '../../input/0';
open(FH, '<', $inputfile) or die $!;

my @input = ();
while(<FH>){
    push(@input, $_)
}

sub part1() {
    return 1
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