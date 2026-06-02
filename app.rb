require 'sinatra/base'
class App < Sinatra::Base
  set :bind, '0.0.0.0'
  set :port, ENV.fetch('PORT', '3000')
  get('/') { 'axhub sinatra ok' }
  run! if app_file == $PROGRAM_NAME
end
