#!/usr/bin/env ruby

class Count
  attr_writer :path
  attr_writer :name
  attr_writer :client_count
  attr_writer :metadata_count

  def print
    padding = 15
    printf "%s\t %s\n", "#{@name} client", @client_count
    printf "%s\t %s\n", "#{@name} metadata", @metadata_count
  end
end

class Counter
  def self.counts(paths)
    new.counts(paths)
  end

  def counts(paths)
    paths.map do |p|
      c = Count.new
      c.path = p
      c.name = p.split("/").last
      c.client_count = %x{wc -l '#{p}/codegen_client.go'}.to_i
      c.metadata_count = %x{wc -l '#{p}/codegen_metadata.go'}.to_i
      c
    end
  end
end

rsc_path = ARGV[0]
if rsc_path == nil
  puts "usage: #{$0} RSC_PATH"
  exit 1
end

# Sanity check
if !File.exist?(File.join(rsc_path, "cm15"))
  puts "invalid argument #{rsc_path}"
  puts "usage: #{$0} RSC_PATH"
  exit 1
end

paths = ["cm15", "cm16", "rl10", "ss/ssc", "ss/ssd", "ss/ssm"].map do |p|
  File.join(rsc_path, p)
end

counts = Counter.counts(paths)
counts.each {|c| c.print }
