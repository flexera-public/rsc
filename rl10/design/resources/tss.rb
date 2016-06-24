module Resources
  class TSS
    include Praxis::ResourceDefinition

    description 'Manipulate the TSS proxy (this is deprecated, please use the /rll/tss/control resource)'
    media_type 'text/plain'

    routing do
      prefix '/rll/tss'
    end

    action :get_hostname do
      description 'Get the TSS hostname to proxy (deprecated, RL10 knows the hostname)'
      routing { get '/hostname' }
      response :ok
    end

    action :put_hostname do
      description 'Set the TSS hostname to proxy (deprecated, RL10 knows the hostname)'
      routing { put '/hostname' }
      params do
        attribute :hostname, String, required: true
      end
      response :ok
    end
  end

  class TSSPlugins
    include Praxis::ResourceDefinition

    description 'TSS Custom Plugins'
    media_type 'text/plain'

    routing do
      prefix '/rll/tss'
    end

    action :index do
      description 'Get TSS plugins list'
      routing { get '/exec' }
      response :ok
    end

    action :create do
      description 'Add new TSS custom plugin'
      routing { put '/exec/:name' }
      params do
        attribute :executable, String, required: true
        attribute :arguments, Attributor::Collection.of(String), required: false
      end
      response :ok
    end

    action :show do
      description 'Get TSS plugin info'
      routing { get '/exec/:name' }
      response :ok
    end

    action :update do
      description 'Update TSS custom plugin'
      routing { put '/exec/:name' }
      params do
        attribute :executable, String, required: true
        attribute :arguments, Attributor::Collection.of(String), required: false
      end
      response :no_content
    end

    action :destroy do
      description 'Delete TSS plugin info'
      routing { delete '/exec/:name' }
      response :no_content
    end

  end

  class TSSControl
    include Praxis::ResourceDefinition

    description 'Manipulate monitoring (TSS) settings'
    media_type 'text/plain'

    routing do
      prefix '/rll/tss/control'
    end

    action :show do
      description 'Show monitoring features'
      routing { get '' }
      response :ok
    end

    action :update do
      description 'Enable/disable monitoring features'
      routing { put '' }
      params do
        attribute :tss_id, String, required: false
        attribute :enable_monitoring, Attributor::String, required: false,
          values:['false','true','none','util','extra','all']
      end
      response :ok
    end

    action :put_control do
      description 'Control the TSS monitoring (deprecated, use the update action)'
      routing { put '' }
      params do
        attribute :tss_id, String, required: false
        attribute :enable_monitoring, Attributor::String, required: false, example: 'all'
      end
      response :ok
    end

  end

end
