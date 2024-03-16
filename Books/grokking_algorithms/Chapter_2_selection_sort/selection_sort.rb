def find_smallest(arr)
  smallest = arr[0]
  index = 0
  arr.each_with_index do |value, i|
    if value < smallest
      smallest = value
      index = i
    end
  end
  index
end

puts find_smallest([1, 2, 3, 4, 5, 6, 7, 8, 9, 10]) # 0
puts find_smallest([10, 9, 8, -7, 6, 5, 4, 3, 2, 1]) # 3
puts find_smallest([1, 10, 2, 9, 3, 8, 4, 7, 5, 6, 0]) # 10

def selection_sort(arr)
  result = []
  arr.length.times do
    smallest = find_smallest(arr)
    result << arr[smallest]
    arr.delete_at(smallest)
  end
  result.inspect
end

puts selection_sort([5, 3, 6, 2, 10]) # [2, 3, 5, 6, 10]
puts selection_sort([10, 9, 8, 7, 6, 5, 4, 3, 2, 1]) # [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]
puts selection_sort([1, 10, 2, 9, 3, 8, 4, 7, 5, 6, 0]) # [0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10]
puts selection_sort([1, 1, 2, 9, 3, 3, 3, 8, 4, 7, 3, 5, 6, 0, 3 ]) # [0, 1, 1, 2, 3, 3, 3, 3, 3, 4, 5, 6, 7, 8, 9]
