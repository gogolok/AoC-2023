lines = File.readlines(ARGV[0], chomp: true)

scratchcardsPerId = (0..lines.length-1).map { 0 }

lines.each_with_index do |line, index|
  (winning_numbers_str, picked_numbers_str) = line.split(':')[1].split('|')

  winning_numbers = winning_numbers_str.split().map(&:to_i)
  picked_numbers = picked_numbers_str.split().map(&:to_i)

  real_winning_numbers = winning_numbers & picked_numbers

  if real_winning_numbers.length < 1
    scratchcardsPerId[index] += 1 # count original card
    next
  end

  n = [scratchcardsPerId.length - index, real_winning_numbers.length].min
  n.times do |idx|
    scratchcardsPerId[index + idx + 1] += 1 + scratchcardsPerId[index]
  end

  scratchcardsPerId[index] += 1 # count original card
end

puts scratchcardsPerId.sum
