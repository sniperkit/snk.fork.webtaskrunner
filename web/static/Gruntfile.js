module.exports = function (grunt) {

    grunt.initConfig({
        copy: {
            vendors: {
                files: [
                    {
                        src: 'node_modules/ansi_up/ansi_up.js',
                        dest: 'vendor/ansi_up/ansi_up.js'
                    },
                    {
                        expand: true,
                        flatten: true,
                        src: 'node_modules/font-awesome/css/*',
                        dest: 'vendor/font-awesome/css/'
                    },
                    {
                        expand: true,
                        flatten: true,
                        src: 'node_modules/font-awesome/fonts/*',
                        dest: 'vendor/font-awesome/fonts/'
                    },
                    {
                        src: 'node_modules/normalize.css/normalize.css',
                        dest: 'vendor/normalize.css/normalize.css'
                    },
                    {
                        src: 'node_modules/vue/dist/vue.min.js',
                        dest: 'vendor/vue/vue.min.js'
                    },
                    {
                        src: 'node_modules/vue-resource/dist/vue-resource.min.js',
                        dest: 'vendor/vue-resource/vue-resource.min.js'
                    }
                ]
            }
        }
    });

    grunt.loadNpmTasks('grunt-contrib-copy');

    grunt.registerTask('default', ['copy']);

};


