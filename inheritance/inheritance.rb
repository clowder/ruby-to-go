class Bird
  def speak
    'Squark'
  end

  def name
    'Generic Bird Type'
  end
end

class Parrot < Bird
  def name
    'Parrot'
  end
end

bird = Parrot.new

p "The #{ bird.name } says #{ bird.speak }."
