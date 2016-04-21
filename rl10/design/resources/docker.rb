module Resources
  class DockerControl
    include Praxis::ResourceDefinition

    description 'Manipulate the Docker integration in RightLink 10'
    media_type 'text/plain'

    routing do
      prefix '/rll/docker/control'
    end

    action :show do
      description 'Show Docker integration features'
      routing { get '' }
      response :ok
    end

    action :update do
      description 'Enable/disable Docker integration features'
      routing { put '' }
      params do
        attribute :enable_docker, Attributor::String, required: false, values: ['none', 'tags', 'all']
        attribute :docker_host, String, required: false
      end
      response :ok
    end
  end
end
