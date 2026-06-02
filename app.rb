require 'sinatra/base'

class DetectApp < Sinatra::Base
  get '/' do
    "language=ruby message=#{ENV.fetch('APPHUB_QA_RUBY_MESSAGE', 'ruby ok')}"
  end
end
