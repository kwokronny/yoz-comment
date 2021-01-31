const gulp = require("gulp");
const stylus = require("gulp-stylus");

gulp.task("css", function () {
  gulp.watch("./demo/demo.styl", function () {
    return gulp.src("./demo/demo.styl").pipe(stylus()).pipe(gulp.dest("./demo"));
  });
});
