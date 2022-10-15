/* eslint-disable no-unused-vars */
const Service = require('./Service');

/**
* unsubscribe from notifications
*
* subscriptionId String Event Subscription ID
* no response value expected for this operation
* */
const deleteIndividualSubcription = ({ subscriptionId }) => new Promise(
  async (resolve, reject) => {
    try {
      resolve(Service.successResponse({
        subscriptionId,
      }));
    } catch (e) {
      reject(Service.rejectResponse(
        e.message || 'Invalid input',
        e.status || 405,
      ));
    }
  },
);
/**
* retrieve subscription
*
* subscriptionId String Event Subscription ID
* suppFeat Object Features supported by the NF service consumer (optional)
* returns NefEventExposureSubsc
* */
const getIndividualSubcription = ({ subscriptionId, suppFeat }) => new Promise(
  async (resolve, reject) => {
    try {
      resolve(Service.successResponse({
        subscriptionId,
        suppFeat,
      }));
    } catch (e) {
      reject(Service.rejectResponse(
        e.message || 'Invalid input',
        e.status || 405,
      ));
    }
  },
);
/**
* update subscription
*
* subscriptionId String Event Subscription ID
* nefEventExposureSubsc NefEventExposureSubsc 
* returns NefEventExposureSubsc
* */
const replaceIndividualSubcription = ({ subscriptionId, nefEventExposureSubsc }) => new Promise(
  async (resolve, reject) => {
    try {
      resolve(Service.successResponse({
        subscriptionId,
        nefEventExposureSubsc,
      }));
    } catch (e) {
      reject(Service.rejectResponse(
        e.message || 'Invalid input',
        e.status || 405,
      ));
    }
  },
);

module.exports = {
  deleteIndividualSubcription,
  getIndividualSubcription,
  replaceIndividualSubcription,
};
