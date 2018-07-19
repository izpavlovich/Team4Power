'use strict';
var dataProvider = require('../data/user.js');
const appInsights = require("express-ai");
/**
 * Operations on /user
 */
module.exports = {
    /**
     * summary: 
     * description: List all user profiles
     * parameters: 
     * produces: 
     * responses: 200, default
     */
    get: function getAllUsers(req, res, next) {
        /**
         * Get the data for response 200
         * For response `default` status 200 is used.
         */
        var status = 200;
        var provider = dataProvider['get']['200'];
        appInsights.setup("91c2e8a3-5944-4ce4-bc6c-e5ee730cb607");
        appInsights.start();
        provider(req, res, function (err, data) {
            if (err) {
                next(err);
                return;
            }
            res.status(status).send(data && data.responses);
        });
    }
};
