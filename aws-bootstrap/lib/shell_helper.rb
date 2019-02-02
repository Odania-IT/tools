class ShellHelper
  def self.execute(cmd)
    IO.popen(cmd) do |lines|
      lines.each(&method(:puts))
    end

    unless $?.success?
      puts 'ERROR: Command Failed!!'
      exit 1
    end
  end
end
