module Resources
  class LoginControl
    include Praxis::ResourceDefinition

    description 'Manipulate login policy settings'
    media_type 'text/plain'

    routing do
      prefix '/rll/login/control'
    end

    action :show do
      description 'Show login policy features'
      routing { get '' }
      response :ok
    end

    action :update do
      description 'Enable/disable login policy features'
      routing { put '' }
      params do
        attribute :enable_login, Attributor::Boolean, required: false
      end
      response :ok
    end

  end

end
