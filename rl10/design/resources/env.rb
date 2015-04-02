module Resources
  class Env
    include Praxis::ResourceDefinition

    description 'Manipulate global script environment variables'
    media_type 'text/plain'

    routing do
      prefix '/rll/env'
    end

    action :index do
      description 'Retrieve all environment variables'
      routing { get '' }
      response :ok
    end

    action :show do
      description 'Retrieve environment variable value'
      routing { get '/:name' }
      response :ok
    end

    action :update do
      description 'Set environment variable value'
      routing { put '/:name' }
      params do
        attribute :name, String, required: true
      end
      payload String, required: true
      response :ok
    end

    action :delete do
      description 'Delete environment variable'
      routing { delete '/:name' }
      response :no_content
    end

  end
end
