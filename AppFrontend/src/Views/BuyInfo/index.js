import React, { Component } from 'react';
import { Text, View } from 'react-native';

export default class BuyInfoScreen extends Component {
    static navigationOptions = {
        headerTitle: (<Text style={{flex:1, color: '#298BFF', fontSize: 25, textAlign: 'center'}}>我的求购信息</Text>)
    };

    render() {
        return (
            <View>
                <Text>
                    这是求购信息界面
                </Text>
            </View>
        )
    }
}