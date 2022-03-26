package controllers

import play.api.libs.json.JsValue
import play.api.mvc.{AbstractController, Action, AnyContent, ControllerComponents}

import javax.inject.{Inject, Singleton}
import scala.collection.mutable.ListBuffer

@Singleton
class Crud1Controller @Inject()(cc: ControllerComponents) extends AbstractController(cc) {
  val crud1List = new ListBuffer[String]()
  crud1List.addAll(List("Crud1 El1", "Crud1 El2", "Crud1 El3", "Crud1 El4", "Crud1 El5", "Crud1 El6"))

  def showAll(): Action[AnyContent] = Action {
    Ok(crud1List.toList.toString())
  }

  def showById(id: Int): Action[AnyContent] = Action {
    if (id < 0 || id >= crud1List.size) {
      NotFound(s"""Element with $id doesn't exist.""")
    } else {
      Ok(crud1List(id))
    }
  }

  def add(): Action[AnyContent] = Action { request =>
    val body: AnyContent = request.body
    val jsonBody: Option[JsValue] = body.asJson

    jsonBody
      .map { json =>
        crud1List += (json \ "el").as[String]
        Ok("Element added")
      }
      .getOrElse {
        BadRequest(s"""Body has to contain json with el element""")
      }
  }

  def update(id: Int): Action[AnyContent] = Action { request =>
    if (id < 0 || id >= crud1List.size) {
      BadRequest(s"""Element with $id doesn't exist.""")
    } else {
      val body: AnyContent = request.body
      val jsonBody: Option[JsValue] = body.asJson

      jsonBody
        .map { json =>
          crud1List.remove(id)
          crud1List.insert(id, (json \ "el").as[String])
          Ok("Element replaced")
        }
        .getOrElse {
          BadRequest(s"""Body has to contain json with el element""")
        }
    }
  }

  def delete(id: Int): Action[AnyContent] = Action {
    crud1List.remove(id)
    Ok("Element removed")
  }

}
