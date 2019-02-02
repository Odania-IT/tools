#!/usr/bin/env ruby
require 'erb'
require 'fileutils'
require 'yaml'

require_relative 'lib/render_template'
require_relative 'lib/shell_helper'

$stdin.sync = true

target_dir = Dir.pwd
base_dir = File.absolute_path File.join(__FILE__ , '..')
config_file = ARGV[0].nil? ? File.join(target_dir, 'config.yml') : ARGV[0]
unless File.exist? config_file
  puts "Config file not found! File: #{config_file}"
  exit 1
end

target_dir = File.dirname config_file
config = YAML.load_file config_file

puts 'Writing accounts-policies.tf'
renderer = RenderTemplate.new File.join(base_dir, 'templates', 'accounts-policies.tf.erb'), File.join(target_dir, 'accounts-policies.tf'), config
renderer.write

puts 'Writing accounts-roles.tf'
renderer = RenderTemplate.new File.join(base_dir, 'templates', 'accounts-roles.tf.erb'), File.join(target_dir, 'accounts-roles.tf'), config
renderer.write

puts 'Writing main.tf'
renderer = RenderTemplate.new File.join(base_dir, 'templates', 'main.tf.erb'), File.join(target_dir, 'main.tf'), config
renderer.write

puts 'Writing main-account-groups.tf'
renderer = RenderTemplate.new File.join(base_dir, 'templates', 'main-account-groups.tf.erb'), File.join(target_dir, 'main-account-groups.tf'), config
renderer.write

puts 'Writing main-account-users.tf'
renderer = RenderTemplate.new File.join(base_dir, 'templates', 'main-account-users.tf.erb'), File.join(target_dir, 'main-account-users.tf'), config
renderer.write

puts 'Writing terraform.tf'
renderer = RenderTemplate.new File.join(base_dir, 'templates', 'terraform.tf.erb'), File.join(target_dir, 'terraform.tf'), config
renderer.write

puts 'Writing variables.tf'
renderer = RenderTemplate.new File.join(base_dir, 'templates', 'variables.tf.erb'), File.join(target_dir, 'variables.tf'), config
renderer.write

# puts 'Copying modules'
# FileUtils.cp_r File.join(base_dir, 'modules'), File.join(target_dir, 'modules')

puts 'Executing plan'
cmd = %{
set -e
cd #{target_dir}
terraform init
terraform plan -input=false -out terraform.plan
}

puts "Terraform plan in folder #{target_dir}"
puts "Executing: #{cmd}"
ShellHelper.execute(cmd)

puts 'Apply? (y/n)'
input = $stdin.gets.chomp
unless 'y'.eql? input
  puts
  puts 'Not applying plan! Continue...'
  puts
  puts
  return
end

cmd = %{
set -e
cd #{target_dir}
terraform apply terraform.plan
}
puts "Applying terraform plan: #{cmd}"
ShellHelper.execute(cmd)
