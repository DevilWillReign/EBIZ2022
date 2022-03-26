package controllers

import play.api.libs.json.JsValue
import play.api.mvc.{AbstractController, Action, AnyContent, ControllerComponents}

import javax.inject.{Inject, Singleton}
import scala.collection.mutable.ListBuffer

@Singleton
class Crud2Controller @Inject()(cc: ControllerComponents) extends AbstractController(cc) {
  val crud2List = new ListBuffer[String]()
  crud2List.addAll(List("Crud2 El1", "Crud2 El2", "Crud2 El3", "Crud2 El4", "Crud2 El5", "Crud2 El6"))

  def showAll(): Action[AnyContent] = Action {
    Ok(crud2List.toList.toString())
  }

  def showById(id: Int): Action[AnyContent] = Action {
    if (id < 0 || id >= crud2List.size) {
      NotFound(s"""Element with $id doesn't exist.""")
    } else {
      Ok(crud2List(id))
    }
  }

  def add(): Action[AnyContent] = Action { request =>
    val body: AnyContent = request.body
    val jsonBody: Option[JsValue] = body.asJson

    jsonBody
      .map { json =>
        crud2List += (json \ "el").as[String]
        Ok("Element added")
      }
      .getOrElse {
        BadRequest(s"""Body has to contain json with el element""")
      }
  }

  def update(id: Int): Action[AnyContent] = Action { request =>
    if (id < 0 || id >= crud2List.size) {
      BadRequest(s"""Element with $id doesn't exist.""")
    } else {
      val body: AnyContent = request.body
      val jsonBody: Option[JsValue] = body.asJson

      jsonBody
        .map { json =>
          crud2List.remove(id)
          crud2List.insert(id, (json \ "el").as[String])
          Ok("Element replaced")
        }
        .getOrElse {
          BadRequest(s"""Body has to contain json with el element""")
        }
    }
  }

  def delete(id: Int): Action[AnyContent] = Action {
    crud2List.remove(id)
    Ok("Element removed")
  }

}
