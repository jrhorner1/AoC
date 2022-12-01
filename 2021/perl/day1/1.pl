#!/usr/bin/perl
# Pass 1 or 2 as a cli argument to solve for part 1 or part 2

use strict;
use warnings;

my $inputfile = '../../input/1';
open(FH, '<', $inputfile) or die $!;

my @input = ();
while(<FH>){
    push(@input, $_)
}

sub linear(){
    my $count = 0;
    my $prev = $input[0];

    for(my $i = 1; $i <= $#input; $i++) {
        if($input[$i] > $prev){
            $count++;
        }
        $prev = $input[$i];
    }
    return $count;
}

sub sliding(){
    my $count = 0;
    my $prev = $input[0] + $input[1] + $input[2];

    for(my $i = 1; $i <= $#input-2; $i++) {
        my $value = $input[$i] + $input[$i+1] + $input[$i+2];
        if($value > $prev){
            $count++;
        }
        $prev = $value
    }
    return $count;
}

if($ARGV[0] == 1){
    print linear(), "\n";
} else {
    print sliding(), "\n";
}
close(FH);