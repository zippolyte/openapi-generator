/**
 * OpenAPI Petstore
 * This is a sample server Petstore server. For this sample, you can use the api key `special-key` to test the authorization filters.
 *
 * The version of the OpenAPI document: 1.0.0
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

#include <QJsonArray>
#include <QJsonDocument>
#include <QJsonObject>
#include <QVariantMap>
#include <QDebug>

#include "OAIHelpers.h"
#include "OAIUserApiRequest.h"

namespace OpenAPI {

OAIUserApiRequest::OAIUserApiRequest(QHttpEngine::Socket *s, QSharedPointer<OAIUserApiHandler> hdl) : QObject(s), socket(s), handler(hdl) {
    auto headers = s->headers();
    for(auto itr = headers.begin(); itr != headers.end(); itr++) {
        requestHeaders.insert(QString(itr.key()), QString(itr.value()));
    }
}

OAIUserApiRequest::~OAIUserApiRequest(){
    disconnect(this, nullptr, nullptr, nullptr);
    qDebug() << "OAIUserApiRequest::~OAIUserApiRequest()";
}

QMap<QString, QString>
OAIUserApiRequest::getRequestHeaders() const {
    return requestHeaders;
}

void OAIUserApiRequest::setResponseHeaders(const QMultiMap<QString, QString>& headers){
    for(auto itr = headers.begin(); itr != headers.end(); ++itr) {
        responseHeaders.insert(itr.key(), itr.value());
    }
}


QHttpEngine::Socket* OAIUserApiRequest::getRawSocket(){
    return socket;
}


void OAIUserApiRequest::createUserRequest(){
    qDebug() << "/v2/user";
    connect(this, &OAIUserApiRequest::createUser, handler.data(), &OAIUserApiHandler::createUser);

    
 
    
    QJsonDocument doc;
    socket->readJson(doc);
    QJsonObject obj = doc.object();
    OAIUser body;
    ::OpenAPI::fromJsonValue(body, obj);
    

    emit createUser(body);
}


void OAIUserApiRequest::createUsersWithArrayInputRequest(){
    qDebug() << "/v2/user/createWithArray";
    connect(this, &OAIUserApiRequest::createUsersWithArrayInput, handler.data(), &OAIUserApiHandler::createUsersWithArrayInput);

    
 
    QJsonDocument doc;
    QList<OAIUser> body;
    if(socket->readJson(doc)){
        QJsonArray jsonArray = doc.array();
        foreach(QJsonValue obj, jsonArray) {
            OAIUser o;
            ::OpenAPI::fromJsonValue(o, obj);
            body.append(o);
        }
    }
    

    emit createUsersWithArrayInput(body);
}


void OAIUserApiRequest::createUsersWithListInputRequest(){
    qDebug() << "/v2/user/createWithList";
    connect(this, &OAIUserApiRequest::createUsersWithListInput, handler.data(), &OAIUserApiHandler::createUsersWithListInput);

    
 
    QJsonDocument doc;
    QList<OAIUser> body;
    if(socket->readJson(doc)){
        QJsonArray jsonArray = doc.array();
        foreach(QJsonValue obj, jsonArray) {
            OAIUser o;
            ::OpenAPI::fromJsonValue(o, obj);
            body.append(o);
        }
    }
    

    emit createUsersWithListInput(body);
}


void OAIUserApiRequest::deleteUserRequest(const QString& usernamestr){
    qDebug() << "/v2/user/{username}";
    connect(this, &OAIUserApiRequest::deleteUser, handler.data(), &OAIUserApiHandler::deleteUser);

    
    QString username;
    fromStringValue(usernamestr, username);
    

    emit deleteUser(username);
}


void OAIUserApiRequest::getUserByNameRequest(const QString& usernamestr){
    qDebug() << "/v2/user/{username}";
    connect(this, &OAIUserApiRequest::getUserByName, handler.data(), &OAIUserApiHandler::getUserByName);

    
    QString username;
    fromStringValue(usernamestr, username);
    

    emit getUserByName(username);
}


void OAIUserApiRequest::loginUserRequest(){
    qDebug() << "/v2/user/login";
    connect(this, &OAIUserApiRequest::loginUser, handler.data(), &OAIUserApiHandler::loginUser);

    
    QString username;
    if(socket->queryString().keys().contains("username")){
        fromStringValue(socket->queryString().value("username"), username);
    }
    
    QString password;
    if(socket->queryString().keys().contains("password")){
        fromStringValue(socket->queryString().value("password"), password);
    }
    


    emit loginUser(username, password);
}


void OAIUserApiRequest::logoutUserRequest(){
    qDebug() << "/v2/user/logout";
    connect(this, &OAIUserApiRequest::logoutUser, handler.data(), &OAIUserApiHandler::logoutUser);

    


    emit logoutUser();
}


void OAIUserApiRequest::updateUserRequest(const QString& usernamestr){
    qDebug() << "/v2/user/{username}";
    connect(this, &OAIUserApiRequest::updateUser, handler.data(), &OAIUserApiHandler::updateUser);

    
    QString username;
    fromStringValue(usernamestr, username);
     
    
    QJsonDocument doc;
    socket->readJson(doc);
    QJsonObject obj = doc.object();
    OAIUser body;
    ::OpenAPI::fromJsonValue(body, obj);
    

    emit updateUser(username, body);
}



void OAIUserApiRequest::createUserResponse(){
    setSocketResponseHeaders();
    socket->setStatusCode(QHttpEngine::Socket::OK);
    socket->writeHeaders();
    if(socket->isOpen()){
        socket->close();
    }
}

void OAIUserApiRequest::createUsersWithArrayInputResponse(){
    setSocketResponseHeaders();
    socket->setStatusCode(QHttpEngine::Socket::OK);
    socket->writeHeaders();
    if(socket->isOpen()){
        socket->close();
    }
}

void OAIUserApiRequest::createUsersWithListInputResponse(){
    setSocketResponseHeaders();
    socket->setStatusCode(QHttpEngine::Socket::OK);
    socket->writeHeaders();
    if(socket->isOpen()){
        socket->close();
    }
}

void OAIUserApiRequest::deleteUserResponse(){
    setSocketResponseHeaders();
    socket->setStatusCode(QHttpEngine::Socket::OK);
    socket->writeHeaders();
    if(socket->isOpen()){
        socket->close();
    }
}

void OAIUserApiRequest::getUserByNameResponse(const OAIUser& res){
    setSocketResponseHeaders();
    QJsonDocument resDoc(::OpenAPI::toJsonValue(res).toObject());
    socket->writeJson(resDoc);
    if(socket->isOpen()){
        socket->close();
    }
}

void OAIUserApiRequest::loginUserResponse(const QString& res){
    setSocketResponseHeaders();
    socket->write(::OpenAPI::toStringValue(res).toUtf8());
    if(socket->isOpen()){
        socket->close();
    }
}

void OAIUserApiRequest::logoutUserResponse(){
    setSocketResponseHeaders();
    socket->setStatusCode(QHttpEngine::Socket::OK);
    socket->writeHeaders();
    if(socket->isOpen()){
        socket->close();
    }
}

void OAIUserApiRequest::updateUserResponse(){
    setSocketResponseHeaders();
    socket->setStatusCode(QHttpEngine::Socket::OK);
    socket->writeHeaders();
    if(socket->isOpen()){
        socket->close();
    }
}


void OAIUserApiRequest::createUserError(QNetworkReply::NetworkError error_type, QString& error_str){
    Q_UNUSED(error_type); // TODO: Remap error_type to QHttpEngine::Socket errors
    setSocketResponseHeaders();
    socket->setStatusCode(QHttpEngine::Socket::NotFound);
    socket->write(error_str.toUtf8());
    if(socket->isOpen()){
        socket->close();
    }
}

void OAIUserApiRequest::createUsersWithArrayInputError(QNetworkReply::NetworkError error_type, QString& error_str){
    Q_UNUSED(error_type); // TODO: Remap error_type to QHttpEngine::Socket errors
    setSocketResponseHeaders();
    socket->setStatusCode(QHttpEngine::Socket::NotFound);
    socket->write(error_str.toUtf8());
    if(socket->isOpen()){
        socket->close();
    }
}

void OAIUserApiRequest::createUsersWithListInputError(QNetworkReply::NetworkError error_type, QString& error_str){
    Q_UNUSED(error_type); // TODO: Remap error_type to QHttpEngine::Socket errors
    setSocketResponseHeaders();
    socket->setStatusCode(QHttpEngine::Socket::NotFound);
    socket->write(error_str.toUtf8());
    if(socket->isOpen()){
        socket->close();
    }
}

void OAIUserApiRequest::deleteUserError(QNetworkReply::NetworkError error_type, QString& error_str){
    Q_UNUSED(error_type); // TODO: Remap error_type to QHttpEngine::Socket errors
    setSocketResponseHeaders();
    socket->setStatusCode(QHttpEngine::Socket::NotFound);
    socket->write(error_str.toUtf8());
    if(socket->isOpen()){
        socket->close();
    }
}

void OAIUserApiRequest::getUserByNameError(const OAIUser& res, QNetworkReply::NetworkError error_type, QString& error_str){
    Q_UNUSED(error_type); // TODO: Remap error_type to QHttpEngine::Socket errors
    setSocketResponseHeaders();
    Q_UNUSED(error_str);  // response will be used instead of error string
    QJsonDocument resDoc(::OpenAPI::toJsonValue(res).toObject());
    socket->writeJson(resDoc);
    if(socket->isOpen()){
        socket->close();
    }
}

void OAIUserApiRequest::loginUserError(const QString& res, QNetworkReply::NetworkError error_type, QString& error_str){
    Q_UNUSED(error_type); // TODO: Remap error_type to QHttpEngine::Socket errors
    setSocketResponseHeaders();
    Q_UNUSED(error_str);  // response will be used instead of error string
    socket->write(::OpenAPI::toStringValue(res).toUtf8());
    if(socket->isOpen()){
        socket->close();
    }
}

void OAIUserApiRequest::logoutUserError(QNetworkReply::NetworkError error_type, QString& error_str){
    Q_UNUSED(error_type); // TODO: Remap error_type to QHttpEngine::Socket errors
    setSocketResponseHeaders();
    socket->setStatusCode(QHttpEngine::Socket::NotFound);
    socket->write(error_str.toUtf8());
    if(socket->isOpen()){
        socket->close();
    }
}

void OAIUserApiRequest::updateUserError(QNetworkReply::NetworkError error_type, QString& error_str){
    Q_UNUSED(error_type); // TODO: Remap error_type to QHttpEngine::Socket errors
    setSocketResponseHeaders();
    socket->setStatusCode(QHttpEngine::Socket::NotFound);
    socket->write(error_str.toUtf8());
    if(socket->isOpen()){
        socket->close();
    }
}


void OAIUserApiRequest::sendCustomResponse(QByteArray & res, QNetworkReply::NetworkError error_type){
    Q_UNUSED(error_type); // TODO
    socket->write(res);
    if(socket->isOpen()){
        socket->close();
    }    
}

void OAIUserApiRequest::sendCustomResponse(QIODevice *res, QNetworkReply::NetworkError error_type){
    Q_UNUSED(error_type);  // TODO
    socket->write(res->readAll());
    if(socket->isOpen()){
        socket->close();
    }
}

}
