/** @type {import('ts-jest').JestConfigWithTsJest} **/
module.exports = {
  setupFilesAfterEnv: ['<rootDir>/jest-setup.ts'],
  testEnvironment: "jsdom",
  transform: {
    "^.+\\.(js|jsx|ts|tsx)$": [
      "babel-jest",
      {
        presets: [
          [
            "next/babel",
            { "preset-react": { "runtime": "automatic" } }
          ]
        ]
      }
    ]
  },
  moduleFileExtensions: ["ts", "tsx", "js", "jsx"],
};
