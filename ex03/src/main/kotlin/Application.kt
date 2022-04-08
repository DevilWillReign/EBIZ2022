package com.example

import com.slack.api.bolt.App
import com.slack.api.bolt.AppConfig
import com.slack.api.bolt.request.Request
import com.slack.api.bolt.request.RequestHeaders
import com.slack.api.bolt.response.Response
import com.slack.api.bolt.util.QueryStringParser
import com.slack.api.bolt.util.SlackRequestParser
import com.slack.api.model.block.Blocks.actions
import com.slack.api.model.block.Blocks.asBlocks
import com.slack.api.model.block.Blocks.section
import com.slack.api.model.block.composition.BlockCompositions.markdownText
import com.slack.api.model.block.composition.BlockCompositions.plainText
import com.slack.api.model.block.element.BlockElements.button
import io.ktor.application.Application
import io.ktor.application.ApplicationCall
import io.ktor.application.call
import io.ktor.features.origin
import io.ktor.http.ContentType
import io.ktor.http.HttpStatusCode
import io.ktor.request.queryString
import io.ktor.request.receiveText
import io.ktor.request.uri
import io.ktor.response.header
import io.ktor.response.respondText
import io.ktor.routing.post
import io.ktor.routing.routing
import io.ktor.util.toMap

val appConfig = AppConfig()
val requestParser = SlackRequestParser(appConfig)
val app = App(appConfig)
val products = mapOf(
    Pair("breakfast", listOf("full english", "pancakes")),
    Pair("drinks", listOf("Coca-Cola", "Pepsi", "Water", "Dr. Pepper")),
    Pair("dinner", listOf("Steak and sides", "Fish and chips", "Full course"))
)

fun Application.module() {
    app.command("/categories") { _, ctx ->
        ctx.ack(asBlocks(
            section { section ->
                section.text(markdownText("Select category"))
            },
            actions { actions ->
                actions
                    .elements(
                        products.keys.map { t ->
                            button { btn ->
                                btn.actionId(t)
                                    .text(plainText("Select category $t"))
                            }
                        }
                    )
            }
        ))
    }

    app.command("/products") { _, ctx ->
        ctx.ack(asBlocks(
            section { section ->
                section.text(markdownText("Select product"))
            },
            actions { actions ->
                actions
                    .elements(
                        products.values.flatMap { l ->
                            l.map { t ->
                                button { btn ->
                                    btn.actionId(t)
                                        .text(plainText("Select product $t"))
                                }
                            }
                        }
                    )
            }
        ))
    }

    for (cat in products.keys) {
        app.blockAction(cat) { _, ctx ->
            ctx.respond(asBlocks(
                section { section ->
                    section.text(markdownText("Selected $cat"))
                },
                section { section ->
                    section.text(markdownText("Select product"))
                },
                actions { actions ->
                    actions
                        .elements(
                            products[cat]?.map { t ->
                                button { btn ->
                                    btn.actionId(t)
                                        .text(plainText("Select product $t"))
                                }
                            }
                        )
                }
            ))
            ctx.ack()
        }
    }

    for (prods in products.values) {
        for (prod in prods) {
            app.blockAction(prod) { _, ctx ->
                ctx.respond(asBlocks(
                    section { section ->
                        section.text(markdownText("Selected $prod"))
                    }
                ))
                ctx.ack()
            }
        }
    }

    routing {
        post("/slack/events") {
            respond(call, app.run(parseRequest(call)))
        }
    }
}

suspend fun parseRequest(call: ApplicationCall): Request<*> {
    val requestBody = call.receiveText()
    val queryString = QueryStringParser.toMap(call.request.queryString())
    val headers = RequestHeaders(call.request.headers.toMap())
    return requestParser.parse(
        SlackRequestParser.HttpRequest.builder()
            .requestUri(call.request.uri)
            .queryString(queryString)
            .requestBody(requestBody)
            .headers(headers)
            .remoteAddress(call.request.origin.remoteHost)
            .build()
    )
}

suspend fun respond(call: ApplicationCall, slackResp: Response) {
    for (header in slackResp.headers) {
        for (value in header.value) {
            call.response.header(header.key, value)
        }
    }
    val status = HttpStatusCode.fromValue(slackResp.statusCode)
    if (slackResp.body != null) {
        call.respondText(slackResp.body, ContentType.parse(slackResp.contentType), status)
    }
}

fun main(args: Array<String>): Unit = io.ktor.server.netty.EngineMain.main(args)