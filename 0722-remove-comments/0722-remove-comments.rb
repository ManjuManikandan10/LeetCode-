def remove_comments(source)
  source.join("\n").gsub(/\/(\*.*?\*\/|\/.*?$)/m, '').gsub(/\A\n+|\n+\z/, '').split(/\n+/)
end