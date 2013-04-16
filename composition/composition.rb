module Bird
  def name
    'Generic bird type'
  end
end

module TalkingBird
  def speak
    'Hello polly'
  end
end

class Parrot
  include Bird
  include TalkingBird

  def name
    'Parrot'
  end
end

parrot = Parrot.new

p "The #{ parrot.name } says #{ parrot.speak }."

