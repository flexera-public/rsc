module Resources
  class Proc
    include Praxis::ResourceDefinition

    description 'List of process variables, such as version, identity, and protocol_version'
    media_type 'text/plain'

    routing do
      prefix '/rll/proc'
    end

    action :index do
      description 'List all process variables'
      routing { get '' }
      response :ok
    end

    action :show do
      description 'Retrieve process variable value'
      routing { get '/:name' }
      response :ok
    end

    action :update do
      description 'Set process variable'
      routing { post '/:name' }
      params do
        attribute :name, String, required: true
      end
      payload String, required: true
      response :ok
    end
  end
end
