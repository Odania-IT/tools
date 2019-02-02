class RenderTemplate
  def initialize(src_file, target_file, data)
    @data = data
    @src_file = src_file
    @target_file = target_file

    @binding = binding
    data.each_pair do |key, val|
      @binding.local_variable_set(key.to_sym, val)
    end
  end

  def write
    puts "Writting #{File.basename(@src_file)} to #{@target_file}"
    template = ERB.new File.read(@src_file), nil, '-'
    File.write @target_file, template.result(@binding)
  end

  def determine_variable_type(val)
    return 'list' if val.is_a? Array
    return 'map' if val.is_a? Hash

    'string'
  end

  def build_account_arns(target)
    result = []
    @data['accounts'].each_pair do |account_name, data|
      result << "arn:aws:iam::#{data['account_id']}:#{target}"
    end

    result.join("\",\n\"")
  end
end
