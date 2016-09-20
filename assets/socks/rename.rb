

fns = Dir.glob("*.png") 

fns.each {|fn| 
  puts fn
  fn =~ /sx_00(..).*png/
  num = $1
  `mv #{fn} #{num}.png`
}
