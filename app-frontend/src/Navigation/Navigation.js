import React, { Component } from 'React';
import { View, Text } from 'react-native';
import {
    createBottomTabNavigator,
    createStackNavigator,
    createAppContainer, NavigationActions
} from "react-navigation";
import {Avatar, Image} from 'react-native-elements';

import HomeScreen from '../Views/Home/Home'
import ReleaseScreen from '../Views/Release/Release'
import ContactScreen from '../Views/Contact/Contact'
import UserScreen from '../Views/User/User'
import MyBuyInfoScreen from "../Views/MyBuyInfo/MyBuyInfo"
import MySellInfoScreen from "../Views/MySellInfo/MySellInfo";
import MyHistoryInfoScreen from "../Views/MyHistoryInfo/MyHistoryInfo";
import UserInfoScreen from "../Views/UserInfo/UserInfo";
import BuyInfoScreen from "../Views/BuyInfo/BuyInfo";
import SellInfoScreen from "../Views/SellInfo/SellInfo";
import SearchScreen from "../Views/Search/Search";
import LoginScreen from "../Views/Login/Login";
import TestPage from '../Views/TestPage/TestPage';

class Test extends Component {
    render() {
        return (
            <View >
                <Text>
                    test page!
                </Text>
            </View>
        )
    }
}

const HomeStack = createStackNavigator({
    Home: { screen: HomeScreen },
    BuyInfo: { screen: BuyInfoScreen },
    SellInfo: { screen: SellInfoScreen },
    Search: { screen: SearchScreen },
    TestPage: { screen: TestPage },
});

const ReleaseStack = createStackNavigator({
    Release: { screen: ReleaseScreen },
    Test: { screen: Test},
});

const ContactStack = createStackNavigator({
    Contact: { screen: ContactScreen },
    Test: { screen: Test},
});

const UserStack = createStackNavigator({
    User: { screen: UserScreen },
    MyBuyInfo: { screen: MyBuyInfoScreen },
    MySellInfo: { screen: MySellInfoScreen },
    MyHistoryInfo: { screen: MyHistoryInfoScreen},
    UserInfo: { screen: UserInfoScreen},
    Login: { screen: LoginScreen },
});

const TabBar = createBottomTabNavigator({
    Home: { screen: HomeStack },
    Release: { screen: ReleaseStack },
    Contact: { screen: ContactStack },
    User: { screen: UserStack },
},{
    defaultNavigationOptions: ({ navigation }) => ({
        tabBarOnPress: (args) => {
            if (args.navigation.state.routeName === 'Release') {
                navigation.reset([NavigationActions.navigate({routeName: 'Release'})], 0);
            } else {
                args.defaultHandler();
            }
        },
        tabBarLabel: ({ focusd, tintColor }) => {
            const { routeName } = navigation.state;
            switch (routeName) {
                case 'Home':
                    return <Text style={{color: tintColor, fontSize: 12, textAlign: 'center'}}>首页</Text>
                case 'Release':
                    return <Text style={{color: tintColor, fontSize: 12, textAlign: 'center'}}>发布</Text>
                case 'Contact':
                    return <Text style={{color: tintColor, fontSize: 12, textAlign: 'center'}}>消息</Text>
                case 'User':
                    return <Text style={{color: tintColor, fontSize: 12, textAlign: 'center'}}>个人</Text>
            }
        },
        tabBarIcon: ({ focused, tintColor }) => {
            const { routeName } = navigation.state;
            let iconRoot;
            if (routeName === 'Home') {
                return (
                    focused? (<Avatar
                        rounded size='small'
                        source={require('../assets/icons/home-focused.png')}
                        overlayContainerStyle={{backgroundColor: 'white'}}
                    />) : (<Avatar
                        rounded size='small'
                        source={require('../assets/icons/home.png')}
                        overlayContainerStyle={{backgroundColor: 'white'}}
                    />)
                ) ;
            }
            else if (routeName === 'Release') {
                return (
                    focused? (<Avatar
                        rounded size='small'
                        source={require('../assets/icons/release-focused.png')}
                        iconStyle={{height: 10}}
                        overlayContainerStyle={{backgroundColor: 'white'}}
                    />) : (<Avatar
                        rounded size='small'
                        source={require('../assets/icons/release.png')}
                        overlayContainerStyle={{backgroundColor: 'white'}}
                    />)
                ) ;
            }
            else if (routeName === 'Contact') {
                return (
                    focused? (<Avatar
                        rounded size='small'
                        source={require('../assets/icons/message-focused.png')}
                        overlayContainerStyle={{backgroundColor: 'white'}}
                    />) : (<Avatar
                        rounded size='small'
                        source={require('../assets/icons/message.png')}
                        overlayContainerStyle={{backgroundColor: 'white'}}
                    />)
                ) ;
            }
            else if (routeName === 'User') {
                return (
                    focused? (<Avatar
                        rounded size='small'
                        source={require('../assets/icons/user-focused.png')}
                        overlayContainerStyle={{backgroundColor: 'white'}}
                    />) : (<Avatar
                        rounded size='small'
                        source={require('../assets/icons/user.png')}
                        overlayContainerStyle={{backgroundColor: 'white'}}
                    />)
                ) ;
            }
        }
    })
});
/*
const AppStack = createStackNavigator({
    Tabs: TabBar,
    TestPage: Test,
}, {
    defaultNavigationOptions: {
        headerStyle: {
            backgroundColor: '#ff9800',
        },
        headerTintColor: '#fff',
    }
});*/

export default createAppContainer(TabBar);

export { Test };
