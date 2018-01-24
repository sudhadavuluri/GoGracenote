// routes.go

package main

func initializeRoutes() {

  router.POST("/reguser", registerUser)

  router.POST("/albumsearch", postGraceMovie)
  
  router.POST("/tracksearch", postTrkSrchAttr)
  
  router.POST("/changelanguage", postChangeLang)
  
}