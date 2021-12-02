using System;
using System.Linq;

namespace day08
{
    class Program
    {
        public static int getSolutionPart1(int[] input)
        {
            return input.Where((num, i) => i > 0 && num > input[i - 1]).Count();
        }

        public static int getSolutionPart2(int[] input)
        {
            return input.Where((num, i) => i > 2 && num > input[i - 3]).Count();
        }

        static void Main(string[] args)
        {
            int[] input = parseInput("../input.txt");

            if ("part2".Equals(Environment.GetEnvironmentVariable("part"))) {
                Console.WriteLine(getSolutionPart2(input));
            } else {
                Console.WriteLine(getSolutionPart1(input));
            }

        }

        static int[] parseInput(string filename) {
            return System.IO.File.ReadLines(filename)
            .Select(line => Int32.Parse(line))
            .ToArray();
        }
    }
}
