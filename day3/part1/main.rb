def symbolPositionsPerLine(lines)
  res = []
  lines.each_with_index do |line, index|
    matches = line.gsub('.', "0").to_enum(:scan, /\D/).map { Regexp.last_match}.map { |m| m.begin(0) }
    res[index] = matches
  end
  res
end

ValueRange = Struct.new(:value, :beginning, :ending, keyword_init: true)

def findNumPositions(line)
  matches = line.to_enum(:scan, /\d+/).map { Regexp.last_match }.map do |m|
    ValueRange.new(value: m[0].to_i, beginning: m.begin(0), ending: m.end(0)-1)
  end
end

def isPartOfEngine(index, val_range, sym_pos_per_line)
  [index-1, index, index+1].each do |idx|
    sym_pos_per_line[idx].each do |sym_pos|
      if (sym_pos >= (val_range[:beginning] - 1)) && (sym_pos <= (val_range[:ending] + 1))
        return true
      end
    end
  end
  return false
end

lines = File.readlines(ARGV[0], chomp: true)

lines.unshift("")
lines.push("")

symPosPerLine = symbolPositionsPerLine(lines)

total = 0

lines.each_with_index do |line, index|
  next if index == 0 || index == lines.length-1 # skip artificially injected lines

  puts line
  findNumPositions(line).each do |val_range|
    if isPartOfEngine(index, val_range, symPosPerLine)
      #puts "part of it #{val_range[:value]}"
      total += val_range[:value]
    end
  end
end

puts total
