lines = File.readlines(ARGV[0], chomp: true)

total = 0

lines.each do |line|
  (winning_numbers_str, picked_numbers_str) = line.split(':')[1].split('|')

  winning_numbers = winning_numbers_str.split().map(&:to_i)
  picked_numbers = picked_numbers_str.split().map(&:to_i)

  real_winning_numbers = winning_numbers & picked_numbers

  if real_winning_numbers.length > 0
    total += (1..real_winning_numbers.length).inject { |sum, _ | 2 * sum }
  end
end

puts total
