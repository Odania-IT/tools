#!/usr/bin/env ruby
require 'fileutils'
require 'yaml'

config_file = File.join(Dir.home, '.bash_aliases.d', 'swamp.yml')
config_file = ARGV[0] unless ARGV[0].nil?
extra_script_dir = File.join(Dir.home, '.swamp_extra_scripts.d')

unless File.exist? config_file
  puts "Config File '#{config_file}' does not exist"
  exit 1
end

swamp_data = YAML.load_file config_file
result = ''
swamp_data.each do |group_name, group|
  extra = ''
  if group['extra_script']
    FileUtils.mkdir_p extra_script_dir unless File.directory? extra_script_dir
    script_file = File.join(extra_script_dir, "#{group_name}_script.sh")
    puts "Updating extra script #{script_file}"
    File.write script_file, group['extra_script']
    FileUtils.chmod 0755, script_file
    extra << " && source #{script_file}"
  end

  result << "# Swamp for #{group_name}\n"
  group['roles'].each do |role|
    target_role = group['target_role']
    target_role = role['target_role'] unless role['target_role'].nil?
    result << "alias swamp-#{role['target_profile']}=\"swamp -account #{role['account']} -mfa-device #{group['mfa_device']} -profile #{group['profile']} "
    result << "-target-profile #{role['target_profile']} -target-role #{target_role} -region #{group['region']} -intermediate-profile session-#{group_name} "
    result << "&& export AWS_PROFILE=#{role['target_profile']} && export AWS_REGION=#{group['region']} #{extra}\"\n"
  end
  result << "\n"
end

File.write '/tmp/swamp-aliases.sh', result

bash_aliases_file = File.join(Dir.home, '.bash_aliases.d', 'swamp-aliases.sh')
tmp_aliases_file = '/tmp/swamp-aliases.sh'

puts `diff #{tmp_aliases_file} #{bash_aliases_file}`

if $?.success?
  puts 'Files are equal'
  exit 0
end

puts "Difference detected"
puts "Copy file? (Enter y to continue)"
input = $stdin.gets.chomp

unless 'y'.eql? input
  puts 'OK, we do not continue'
  exit 1
end

puts 'Copying file'
FileUtils.mkdir_p File.base_dir(bash_aliases_file)
FileUtils.cp tmp_aliases_file, bash_aliases_file

puts
puts 'Setup:'
puts 'Make sure you add something like "source ~/.bash_aliases.d/swamp-aliases.sh" to your .bashrc file'

