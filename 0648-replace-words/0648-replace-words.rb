# @param {String[]} dictionary
# @param {String} sentence
# @return {String}
def replace_words(dictionary, sentence)
  # sentence.gsub /\b((#{dictionary.sort_by(&:size)*?|})\S+)/, '\2'
  dictionary.inject(sentence.split) {  |arr, w| arr.map { |ww| ww.start_with?(w) ? w : ww } }.join ' '    
end