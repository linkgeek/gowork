package main

import (
   "github.com/gin-gonic/gin"
   "gopkg.in/go-playground/validator.v9"
   "net/http"
   "time"
)

type Booking struct {
   CheckIn time.Time `form:"check_in" validate:"required,bookableDate" time_format:"2006-01-02"`
   CheckOut time.Time `form:"check_out" validate:"required,gtfield=CheckIn" time_format:"2006-01-02"`
}


func main(){

   r := gin.Default()

   validate := validator.New()
   validate.RegisterValidation("bookableDate", bookableDate)
   r.GET("/bookable", func(c *gin.Context) {
      var book Booking
      if err := c.ShouldBind(&book); err != nil {
         c.JSON(http.StatusInternalServerError, gin.H{
            "error": err.Error(),
         })
         c.Abort()
         return
      }
      if err := validate.Struct(book); err != nil {
         c.JSON(http.StatusInternalServerError, gin.H{
            "error": err.Error(),
         })
         c.Abort()
         return
      }

      c.JSON(http.StatusOK, gin.H{
         "message": "OK",
         "booking": book,
      })
   })

   r.Run()
}

func bookableDate(fl validator.FieldLevel) bool {

   if date, ok := fl.Field().Interface().(time.Time); ok {
      today := time.Now()
      if date.Unix() > today.Unix() {
         return true
      }
   }

   return false
}