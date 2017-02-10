module.exports = function(grunt) {

  // Project configuration.
  grunt.initConfig({
    pkg: grunt.file.readJSON('package.json'),
    exec: {
      one_two_three:{
       cmd: 'echo "1" && sleep 1 && echo "2" && sleep 1 && echo "3" && sleep 1'
       },
      three_two_one:{
       cmd: 'echo "3" && sleep 1 && echo "2" && sleep 1 && echo "1" && sleep 1'
       }
      }

  });

     grunt.loadNpmTasks('grunt-exec');

  // Default task(s).
  grunt.registerTask('test123', ['exec:one_two_three']);
  grunt.registerTask('test321_a_bit_longer_name', ['exec:three_two_one']);

};