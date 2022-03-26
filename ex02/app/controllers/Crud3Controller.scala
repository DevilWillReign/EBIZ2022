package controllers

import play.api.libs.json.JsValue
import play.api.mvc.{AbstractController, Action, AnyContent, ControllerComponents}

import javax.inject.{Inject, Singleton}
import scala.collection.mutable.ListBuffer

@Singleton
class Crud3Controller @Inject()(cc: ControllerComponents) extends AbstractController(cc) {
  val crud3List = new ListBuffer[String]()
  crud3List.addAll(List("Crud3 El1", "Crud3 El2", "Crud3 El3", "Crud3 El4", "Crud3 El5", "Crud3 El6"))

  def showAll(): Action[AnyContent] = Action {
    Ok(crud3List.toList.toString())
  }

  def showById(id: Int): Action[AnyContent] = Action {
    if (id < 0 || id >= crud3List.size) {
      NotFound(s"""Element with $id doesn't exist.""")
    } else {
      Ok(crud3List(id))
    }
  }

  def add(): Action[AnyContent] = Action { request =>
    val body: AnyContent = request.body
    val jsonBody: Option[JsValue] = body.asJson

    jsonBody
      .map { json =>
        crud3List += (json \ "el").as[String]
        Ok("Element added")
      }
      .getOrElse {
        BadRequest(s"""Body has to contain json with el element""")
      }
  }

  def update(id: Int): Action[AnyContent] = Action { request =>
    if (id < 0 || id >= crud3List.size) {
      BadRequest(s"""Element with $id doesn't exist.""")
    } else {
      val body: AnyContent = request.body
      val jsonBody: Option[JsValue] = body.asJson

      jsonBody
        .map { json =>
          crud3List.remove(id)
          crud3List.insert(id, (json \ "el").as[String])
          Ok("Element replaced")
        }
        .getOrElse {
          BadRequest(s"""Body has to contain json with el element""")
        }
    }
  }

  def delete(id: Int): Action[AnyContent] = Action {
    crud3List.remove(id)
    Ok("Element removed")
  }

}
