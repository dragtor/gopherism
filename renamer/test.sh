#!/bin/sh

mkdir -p samples
cd samples
touch birthday_{0001..0003}.txt
#touch 'christmas 2016 ({0001..0003} of 100).txt'
mkdir -p nested
cd nested
touch n_{0001..0003}.txt
