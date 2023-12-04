ValueRange = Struct.new(:value, :beginning, :ending, keyword_init: true)

def numberPositionsPerLine(lines)
  [].tap do |res|
    lines.each_with_index do |line, index|
      res[index] = line.to_enum(:scan, /\d+/).map { Regexp.last_match }.map do |m|
        ValueRange.new(value: m[0].to_i, beginning: m.begin(0), ending: m.end(0)-1)
      end
    end
  end
end

def findSymPositions(line)
  line.to_enum(:scan, /\*/).map { Regexp.last_match}.map { |m| m.begin(0) }
end

def numsNextToSymbol(index, sym_pos, num_pos_per_line)
  [].tap do |nums_next_to_symbol|
    [index-1, index, index+1].each do |idx|
      num_pos_per_line[idx].each do |num_pos|
        if (sym_pos >= (num_pos[:beginning] - 1)) && (sym_pos <= (num_pos[:ending] + 1))
          nums_next_to_symbol << num_pos
        end
      end
    end
  end
end

lines = File.readlines(ARGV[0], chomp: true)
lines.unshift("")
lines.push("")

numPosPerLine = numberPositionsPerLine(lines)
total = 0
lines.each_with_index do |line, index|
  next if index == 0 || index == lines.length-1 # skip artificially injected lines

  findSymPositions(line).each do |sym_pos|
    nums_next_to_symbol = numsNextToSymbol(index, sym_pos, numPosPerLine)
    if nums_next_to_symbol.length == 2
      total += nums_next_to_symbol[0][:value] * nums_next_to_symbol[1][:value]
    end
  end
end

puts total
