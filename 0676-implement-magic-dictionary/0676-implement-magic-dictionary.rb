class MagicDictionary
    def initialize()

    end


=begin
    :type dictionary: String[]
    :rtype: Void
=end
    def build_dict(dictionary)
      @arr = dictionary.each_with_object([]) { _2 << [_1, _1.size] }             
    end


=begin
    :type search_word: String
    :rtype: Boolean
=end
    def search(search_word, sz = search_word.size)
      @arr.select { _1[1] == sz }.any? { |(w, _)| sz.times.inject(0) { |c, i| search_word[i] == w[i] ? c : c == 1 ? (break 2) : 1} == 1 } 
    end


end
