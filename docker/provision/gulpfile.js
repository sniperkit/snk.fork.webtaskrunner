// including plugins
var gulp = require('gulp');
var exec = require('child_process').exec;

gulp.task('test123', function (cb) {
    exec('echo "1" && sleep 1 && echo "2" && sleep 1 && echo "3" && sleep 1', function (err, stdout, stderr) {
        console.log(stdout);
        console.log(stderr);
        cb(err);
    });
})

gulp.task('test321',['test123'], function (cb) {
    exec('echo "3" && sleep 1 && echo "2" && sleep 1 && echo "1" && sleep 1', function (err, stdout, stderr) {
        console.log(stdout);
        console.log(stderr);
        cb(err);
    });
})