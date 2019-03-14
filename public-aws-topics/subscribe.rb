#!/usr/bin/env ruby
require 'yaml'

base_dir = File.absolute_path(File.dirname(__FILE__))
email = ARGV[0]
service = ARGV[1]

if email.nil?
  puts 'No email provided!'
  exit 1
end

config = YAML.load_file(File.join(base_dir, 'services.yml'))

subscribe_to = config['services']
unless service.nil?
  if config['services'][service].nil?
    puts "Could not find service #{service}!"
    puts "Services:\n- #{service_arns.keys.join("\n- ")}"
    exit 2
  end

  subscribe_to = {
      "#{service}" => config['services'][service],
  }
end

subscribe_to.each do |serviceName, topicArn|
  puts "Subscribing to: #{serviceName}"
  puts `aws sns subscribe --region 'us-east-1' --topic-arn "#{topicArn}" --protocol email --notification-endpoint "#{email}"`
end
