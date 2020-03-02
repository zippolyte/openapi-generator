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

#ifndef PFX_PFXPetApi_H
#define PFX_PFXPetApi_H

#include "PFXHttpRequest.h"

#include "PFXApiResponse.h"
#include "PFXHttpFileElement.h"
#include "PFXPet.h"
#include <QString>

#include <QObject>

namespace test_namespace {

class PFXPetApi : public QObject {
    Q_OBJECT

public:
    PFXPetApi(const QString &scheme = "http", const QString &host = "petstore.swagger.io", int port = 0, const QString &basePath = "/v2", const int timeOut = 0);
    ~PFXPetApi();

    void setScheme(const QString &scheme);
    void setHost(const QString &host);
    void setPort(int port);
    void setBasePath(const QString &basePath);
    void setTimeOut(const int timeOut);
    void setWorkingDirectory(const QString &path);
    void addHeaders(const QString &key, const QString &value);
    void enableRequestCompression();
    void enableResponseCompression();
    void abortRequests();

    void addPet(const PFXPet &body);
    void deletePet(const qint64 &pet_id, const QString &api_key);
    void findPetsByStatus(const QList<QString> &status);
    void findPetsByTags(const QList<QString> &tags);
    void getPetById(const qint64 &pet_id);
    void updatePet(const PFXPet &body);
    void updatePetWithForm(const qint64 &pet_id, const QString &name, const QString &status);
    void uploadFile(const qint64 &pet_id, const QString &additional_metadata, const PFXHttpFileElement &file);

private:
    QString _scheme, _host;
    int _port;
    QString _basePath;
    int _timeOut;
    QString _workingDirectory;
    QMap<QString, QString> defaultHeaders;
    bool isResponseCompressionEnabled;
    bool isRequestCompressionEnabled;

    void addPetCallback(PFXHttpRequestWorker *worker);
    void deletePetCallback(PFXHttpRequestWorker *worker);
    void findPetsByStatusCallback(PFXHttpRequestWorker *worker);
    void findPetsByTagsCallback(PFXHttpRequestWorker *worker);
    void getPetByIdCallback(PFXHttpRequestWorker *worker);
    void updatePetCallback(PFXHttpRequestWorker *worker);
    void updatePetWithFormCallback(PFXHttpRequestWorker *worker);
    void uploadFileCallback(PFXHttpRequestWorker *worker);

signals:

    void addPetSignal();
    void deletePetSignal();
    void findPetsByStatusSignal(QList<PFXPet> summary);
    void findPetsByTagsSignal(QList<PFXPet> summary);
    void getPetByIdSignal(PFXPet summary);
    void updatePetSignal();
    void updatePetWithFormSignal();
    void uploadFileSignal(PFXApiResponse summary);

    void addPetSignalFull(PFXHttpRequestWorker *worker);
    void deletePetSignalFull(PFXHttpRequestWorker *worker);
    void findPetsByStatusSignalFull(PFXHttpRequestWorker *worker, QList<PFXPet> summary);
    void findPetsByTagsSignalFull(PFXHttpRequestWorker *worker, QList<PFXPet> summary);
    void getPetByIdSignalFull(PFXHttpRequestWorker *worker, PFXPet summary);
    void updatePetSignalFull(PFXHttpRequestWorker *worker);
    void updatePetWithFormSignalFull(PFXHttpRequestWorker *worker);
    void uploadFileSignalFull(PFXHttpRequestWorker *worker, PFXApiResponse summary);

    void addPetSignalE(QNetworkReply::NetworkError error_type, QString error_str);
    void deletePetSignalE(QNetworkReply::NetworkError error_type, QString error_str);
    void findPetsByStatusSignalE(QList<PFXPet> summary, QNetworkReply::NetworkError error_type, QString error_str);
    void findPetsByTagsSignalE(QList<PFXPet> summary, QNetworkReply::NetworkError error_type, QString error_str);
    void getPetByIdSignalE(PFXPet summary, QNetworkReply::NetworkError error_type, QString error_str);
    void updatePetSignalE(QNetworkReply::NetworkError error_type, QString error_str);
    void updatePetWithFormSignalE(QNetworkReply::NetworkError error_type, QString error_str);
    void uploadFileSignalE(PFXApiResponse summary, QNetworkReply::NetworkError error_type, QString error_str);

    void addPetSignalEFull(PFXHttpRequestWorker *worker, QNetworkReply::NetworkError error_type, QString error_str);
    void deletePetSignalEFull(PFXHttpRequestWorker *worker, QNetworkReply::NetworkError error_type, QString error_str);
    void findPetsByStatusSignalEFull(PFXHttpRequestWorker *worker, QNetworkReply::NetworkError error_type, QString error_str);
    void findPetsByTagsSignalEFull(PFXHttpRequestWorker *worker, QNetworkReply::NetworkError error_type, QString error_str);
    void getPetByIdSignalEFull(PFXHttpRequestWorker *worker, QNetworkReply::NetworkError error_type, QString error_str);
    void updatePetSignalEFull(PFXHttpRequestWorker *worker, QNetworkReply::NetworkError error_type, QString error_str);
    void updatePetWithFormSignalEFull(PFXHttpRequestWorker *worker, QNetworkReply::NetworkError error_type, QString error_str);
    void uploadFileSignalEFull(PFXHttpRequestWorker *worker, QNetworkReply::NetworkError error_type, QString error_str);

    void abortRequestsSignal(); 
};

} // namespace test_namespace
#endif
