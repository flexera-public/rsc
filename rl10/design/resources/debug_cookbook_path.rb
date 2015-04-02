module Resources
  class DebugCookbookPath
    include Praxis::ResourceDefinition

    description 'Manipulate debug cookbook directory location'
    media_type 'text/plain'

    routing do
      prefix '/rll/debug/cookbook'
    end

    action :show do
      description 'Retrieve debug cookbook directory location'
      routing { get '' }
      response :ok
    end

    action :update do
      description 'Set debug cookbook directory location'
      routing { put '' }
      params do
        attribute :path, String, required: true
      end
      response :ok
    end

    action :delete do
      description 'Remove debug cookbook directory location'
      routing { delete '' }
      response :no_content
    end

  end
end
