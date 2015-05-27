module Resources
  class Rl10
    include Praxis::ResourceDefinition

    description 'Miscellaneous RightLink 10 local requests'
    media_type 'text/plain'

    routing do
      prefix '/rll'
    end

    action :upgrade do
      description 'Relaunch the RightLink process using a specified binary'
      routing { post '/upgrade' }
      params do
        attribute :exec, String, required: true, description: 'Absolute path to binary'
      end
      response :ok
    end

    action :run_recipe do
      description 'Run git-based scripts (as recipes) synchronously'
      routing { post '/run/recipe' }
      params do
        attribute :recipe, String, required: true, description: 'Name of recipe'
        attribute :json, String, description: 'JSON hash of "name": "value" pairs'
        attribute :arguments, Attributor::Hash.of(key: String, value: String), description: 'Script argument values'
      end
      response :ok
    end

    action :run_right_script do
      description 'Run RightScripts synchronously'
      routing { post '/run/right_script' }
      params do
        attribute :right_script, String, description: 'Name of script'
        attribute :right_script_id, Integer, description: 'Id of script'
        attribute :arguments, Attributor::Hash.of(key: String, value: String), description: 'Script argument values'
      end
      response :ok
    end

  end
end

