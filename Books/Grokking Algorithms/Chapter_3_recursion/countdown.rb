def countdown(i)
  puts i
  countdown(i - 1) if i.positive?
end

countdown(5)
puts '---'
countdown(10)
puts '---'
countdown(20)
